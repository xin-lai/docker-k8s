# 使用官方镜像
FROM python:3.7-slim

# 设置工作目录
WORKDIR /app

# 复制当前目录
COPY . /app

RUN apt-get install python3-flask

# 安装模块
RUN pip install --trusted-host pypi.python.org -r requirements.txt

# Run app.py when the container launches
CMD ["python", "app.py"]