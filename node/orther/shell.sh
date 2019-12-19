#!/bin/sh
cd ../../src/Hexo.blog
npm install hexo --save
hexo generate
npm install hexo-server --save
hexo server -p 8000 -s -l
