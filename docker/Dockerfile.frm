ARG ALLINONE=default
ARG EDGE_CLOUD_BASE_IMAGE=scratch

FROM $ALLINONE as allinone

FROM $EDGE_CLOUD_BASE_IMAGE

COPY --from=allinone /usr/local/bin/frm /usr/local/bin/frm
COPY --from=allinone /plugins/platforms.so /plugins/
