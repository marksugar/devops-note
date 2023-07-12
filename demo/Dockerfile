FROM registry.cn-hangzhou.aliyuncs.com/marksugar/jdk:8u202
MAINTAINER by mark

ENV JAVA_OPTS="\
    -server \
    -Xms2048m \
    -Xmx2048m \
    -Xmn512m \
    -Xss256k \
    -XX:+UseConcMarkSweepGC \
    -XX:+UseCMSInitiatingOccupancyOnly \
    -XX:CMSInitiatingOccupancyFraction=70 \
    -XX:+HeapDumpOnOutOfMemoryError \
    -XX:HeapDumpPath=/data/logs" \
    MY_USER=linuxea \
    MY_USER_ID=316


RUN addgroup -g ${MY_USER_ID} -S ${MY_USER} \
    && adduser -u ${MY_USER_ID} -S -H -s /sbin/nologin -g 'java' -G ${MY_USER} ${MY_USER} \
    && mkdir /data/logs -p
COPY target/*.jar /data/
#COPY skywalking-agent /
WORKDIR /data
USER linuxea


#CMD java ${JAVA_OPTS} -javaagent:/skywalking-agent/skywalking-agent.jar -jar *.jar
CMD java ${JAVA_OPTS}  -jar *.jar