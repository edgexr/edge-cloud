// Copyright 2022 MobiledgeX, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloudcommon

import (
	"testing"

	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

func TestDecodeK8S(t *testing.T) {
	objs, _, err := DecodeK8SYaml(testKubeManifest)
	require.Nil(t, err)
	require.Equal(t, 4, len(objs))
	_, ok := objs[0].(*v1.Service)
	require.True(t, ok)
	_, ok = objs[1].(*v1.Service)
	require.True(t, ok)
	_, ok = objs[2].(*appsv1.Deployment)
	require.True(t, ok)

}

var testKubeManifest = `apiVersion: v1
kind: Service
metadata:
  name: testapp-tcp-service
  labels:
    run: testapp
spec:
  type: LoadBalancer
  ports:
  - port: 27272
    targetPort: 27272
    protocol: TCP
    name: grpc27272
  - port: 27273
    targetPort: 27273
    protocol: TCP
    name: rest27273
  - port: 27274
    targetPort: 27274
    protocol: TCP
    name: http27274
  - port: 27275
    targetPort: 27275
    protocol: TCP
    name: tcp27275
  selector:
    run: testapp
---
apiVersion: v1
kind: Service
metadata:
  name: testapp-udp-service
  labels:
    run: testapp
spec:
  type: LoadBalancer
  ports:
  - port: 27276
    targetPort: 27276
    protocol: UDP
    name: udp27276
  selector:
    run: testapp
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: testapp-deployment
spec:
  selector:
    matchLabels:
      run: testapp
  replicas: 2
  template:
    metadata:
      labels:
        run: testapp
    spec:
      volumes:
      imagePullSecrets: 
      - name: mexregistrysecret
      containers:
      - name: testapp
        image: registry.mobiledgex.net:5000/mobiledgex/mexexample
        imagePullPolicy: Always
        ports:
        - containerPort: 27272
          protocol: TCP
        - containerPort: 27273
          protocol: TCP
        - containerPort: 27274
          protocol: TCP
        - containerPort: 27275
          protocol: TCP
        - containerPort: 27276
          protocol: UDP
# adding nested 'yaml document separator' inside mqtt_params_robot.yaml below
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: mqtt-bridge-cfg-files-configmap
data:
  mqtt_params_robot.yaml: |
    ---
    mqtt:
      client:
        protocol: $(arg mqtt_protocol)      # MQTTv311
      connection:
        host: $(arg mqtt_host)
        port: $(arg mqtt_port)
        keepalive: $(arg mqtt_keepalive)
      private_path: device/001
    serializer: json:dumps
    deserializer: json:loads
    `

var testDockerComposeManifest = `version: '3.3'

services:
   db:
     image: mysql:5.7
     restart: always
     environment:
       MYSQL_ROOT_PASSWORD: somewordpress
       MYSQL_DATABASE: wordpress
       MYSQL_USER: wordpress
       MYSQL_PASSWORD: wordpress

   wordpress:
     depends_on:
       - db
     image: wordpress:latest
     ports:
       - "8000:80"
     restart: always
     environment:
       WORDPRESS_DB_HOST: db:3306
       WORDPRESS_DB_USER: wordpress
       WORDPRESS_DB_PASSWORD: wordpress
       WORDPRESS_DB_NAME: wordpress
`

func TestDecodeDockerCompose(t *testing.T) {
	containers, err := DecodeDockerComposeYaml(testDockerComposeManifest)
	require.Nil(t, err)
	require.Equal(t, 2, len(containers))
	dbContainer, ok := containers["db"]
	require.True(t, ok, "Container 'db' exists")
	require.Equal(t, dbContainer.Image, "mysql:5.7", "DB container image exists")
	wpContainer, ok := containers["wordpress"]
	require.True(t, ok, "Container 'wordpress' exists")
	require.Equal(t, wpContainer.Image, "wordpress:latest", "Wordpress container image exists")
}
