# docker build -t leo/show-info
FROM busybox:latest

# 先在本机编译生成show-info程序
# GOOS=linux GOARCH=i386 go build -ldflags '-s -w'
COPY ./show-info / 
CMD /show-info