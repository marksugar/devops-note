FROM uhub.service.ucloud.cn/marksugar-k8s/alpine:3.17.4
COPY kustimoze /sbin/kustimoze
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories  && \
    apk fix && \
    apk --no-cache --update add git git-lfs gpg less openssh patch curl yq jq && \
    git lfs install 
ENTRYPOINT ["git"]
CMD ["--help"]



FROM uhub.service.ucloud.cn/marksugar-k8s/alpine:3.17.4
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories  && \
    apk fix && \
    apk --no-cache --update add git git-lfs gpg less openssh patch && \
    git lfs install 
ENTRYPOINT ["git"]
CMD ["--help"]



