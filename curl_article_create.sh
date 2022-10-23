#!/bin/bash

curl --location --request POST 'http://laughing.plus/blog/api/article' \
--form 'title="分别使用Go、Java、TypeScript实现二分查找"' \
--form 'author="yangjiafeng"' \
--form 'category="web"' \
--form 'categoryName="web技术"' \
--form 'imageUrl=@"/Users/yangjiafeng/Desktop/yang/binary.jpeg"'

