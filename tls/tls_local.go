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

package tls

import (
	"crypto/tls"
	"fmt"
)

const LocalTestCert = `-----BEGIN CERTIFICATE-----
MIICpjCCAY4CCQD1NBBxj1nCFTANBgkqhkiG9w0BAQsFADAUMRIwEAYDVQQDDAkx
MjcuMC4wLjEwIBcNMTkxMDAyMDcxNDAwWhgPMjExOTA5MDgwNzE0MDBaMBQxEjAQ
BgNVBAMMCTEyNy4wLjAuMTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
AMNcg7og9fJc0MDGRfR2hHOQUc2gqHr5tTNEDQjRPtMukITB9dJUxGaTOArIiIdx
+g6A58pwxh+yTtR62lIXKvARe1yBwkzqNUZtColBCwmyvGm4OCyRuQCpgOuzwLoC
uGx8M98cGS5NAeZguNIKUV5vYZ/lVnQ8TfKlQg2UZuVvbnUUJGf6h6OY/5JRpeRk
ESmZkp54vUnO8G0OsFGi/AIFA4w7xxk309WEnTbRwmQkmLdIw7VgizclRB9eAv+T
IKMagZ8PM/z2eWz4GXt0Xk1meVnTeCZAMC6BggbCCdSjCMcY7WSG9c8oA+eeTWKb
rlXAnRfG4fqHnWdDEyeTCOUCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAHzdJT/nB
4+bVE4Ds2Gry7iatB5yJ/mp8kstySUmSzSM8GPIUar6c6g6Hu97OKCDEi8VI7KvE
bMSyXSejQh1HRuHg9zxkTvE1YtMz9Vx26tIBN2GfdHGZ9tkhOq4wvUjdoAZhQ2OA
phVvxIWKsjGhB3iY9RzhNVZt0QLh5wV/4BBw+dR4SWg7e2JJoLO0YkEb/Aw4a+3g
h3n5d5xDgLZtG/n4OJehHFjsTfZ3qgHBP8zioSBYURDVJwdQMP+lLPQnnA+jauSf
3u+mpgmit6M+v0kHIfHdYrIL/0o0Qldj5aMR73kKHqxYnWml1MOkFDQMlYgKAYfS
+5GYolB2RNzhlg==
-----END CERTIFICATE-----`

const LocalTestKey = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDDXIO6IPXyXNDA
xkX0doRzkFHNoKh6+bUzRA0I0T7TLpCEwfXSVMRmkzgKyIiHcfoOgOfKcMYfsk7U
etpSFyrwEXtcgcJM6jVGbQqJQQsJsrxpuDgskbkAqYDrs8C6ArhsfDPfHBkuTQHm
YLjSClFeb2Gf5VZ0PE3ypUINlGblb251FCRn+oejmP+SUaXkZBEpmZKeeL1JzvBt
DrBRovwCBQOMO8cZN9PVhJ020cJkJJi3SMO1YIs3JUQfXgL/kyCjGoGfDzP89nls
+Bl7dF5NZnlZ03gmQDAugYIGwgnUowjHGO1khvXPKAPnnk1im65VwJ0XxuH6h51n
QxMnkwjlAgMBAAECggEBAKXhRRxi5lLPSpWVBUso4xg5H6yI0OXyZ1BLZQHqofyw
2hdPlgns3gL22CwX0A2p0phEDgyPqzq4rUSvt/biIxxKy8NljnS+nVPaPIsEhnnK
oT7nJBN7/Gk6g/Qe4fNPYgcDgwvbkOoGiylxUR1Pk32q92rMUZdKqSjx918Odjwy
4Jc3O7yhhiDdUdOz2ELIAivVgbpalez2Pohcn3y8pqyv3fWuYboDawyAFHC60L9n
K4wIKUUCYIckI+u6DKOiuwkcE7jOhHYVZ1ZTGCAHNtXcxKoEAFQql61vJGjnWaUX
s6/M3C9aMipmpjiWAV+acTSK10fo7h1Mp04Ny6YJ0rUCgYEA6k6f6acbQa/v+M8i
+oUSeB8IGx3exuS7hFHsKvyMjuZumMvZiv4MJhzXMqyJbGALz/pBkHo/hd+z299g
GgoMlOIMvwGdsLbDIYt/Jz8WcXQOubfn/XruUnNmrHAT/QCG55vkta7WiYx++//z
PAP6Lh+XLuJiWKA2yU4NCdzg9P8CgYEA1XLTl4ycL7KfOD/G9k94EkgcEN8+nEnQ
HaI69FNCtm6zdN6y8C73MSkcfAzefxG4FCqD2hehDYqwTdI1FfnIq1wbJfKvApQ/
WAd6rIZkDIp+RlRUn/3hIiTKkRjjC2VSHwb6SgJeCTqVMOOa4jxHwtuoTByA4TE6
+nHDpeVMzhsCgYAerNR6vonu+52rMVMfATT9zPI+upZj66YxGJiWuCGew8RO9MEy
VTg59SSnWbdUHO7u95CF3btbR4JAkun9/rrO8xlAGHIQpJy/U4f+F2BU0iDrZ4xT
vzhcSwdyI+o6AUuoHty+fHlR4LbvPv8VKGAkkbY1SSYe/Dqv+Cv00poQ+QKBgEFu
c8KPTgmpHfnmQx2aetpP1JqvdMgAMJYE90GSjG7UPJrk8cGDBC/DxIGaiFI3olA8
QDMI70vxad5RRAi6i3NMRt45KMEHpOtdT5O3ls/pXJA+pbpt3yfSU25rTJ5fru7n
+q3ZT+5QUQ+tlfsoM29eWpcKXc+qgcPoP+uJASRlAoGBAIDiLNMSJy0109obIJIo
EWuGeuFit8zFTxLsFXIuscMbqZ7vV9XA2qTXlrc5fBlTZ0mO9RWkK1z1dRoluU20
Ty5Scst9RNo9U91rAJOAG2Vrgg9058zCOngxKTnT6EygJx8ZFEVMy0iowri7hNq2
U5CcqvvOrDIFqFt/iZALVwjU
-----END PRIVATE KEY-----`

func GetLocalTLSConfig() (*tls.Config, error) {
	certificate, err := tls.X509KeyPair([]byte(LocalTestCert), []byte(LocalTestKey))
	if err != nil {
		return nil, fmt.Errorf("could not load server key pair: %s", err)
	}
	return &tls.Config{
		Certificates:       []tls.Certificate{certificate},
		InsecureSkipVerify: true,
	}, nil
}
