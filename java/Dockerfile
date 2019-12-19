# 基于Java 9
FROM java:openjdk-8

# 设置工作目录
WORKDIR /app

# 复制文件到工作目录
COPY . /app

# 设置Java环境变量
ENV PATH=$PATH:$JAVA_HOME/bin
ENV JRE_HOME=${JAVA_HOME}/jre
ENV CLASSPATH=.:${JAVA_HOME}/lib:${JRE_HOME}/lib

# 编译
RUN ["/usr/lib/jvm/java-8-openjdk-amd64/bin/javac","Hello.java"]

# 运行
ENTRYPOINT ["/usr/lib/jvm/java-8-openjdk-amd64/bin/java", "Hello"]