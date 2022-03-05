FROM golang:1.16-alpine
LABEL cloud.learn=docker
RUN mkdir -p /appHome \
&& echo "emo:x:600:600:emo:/home/emo:/bin/bash" >> /etc/passwd \
&& echo "emo:x:600:" >> /etc/group && mkdir /home/emo
COPY [--chown=emo:emo] ./httpserver /appHome/
USER emo
EXPOSE 8080
WORKDIR /appHome
ENTRYPOINT ["./httpserver"]
