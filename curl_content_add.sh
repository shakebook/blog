#!/bin/bash
code=cat /Users/yangjiafeng/go-src/blog/codeArticle.txt
curl --location --request POST 'http://laughing.plus/blog/api/article/content' \
--header 'Content-Type: application/json' \
--data-raw '{
  "id":12,
  "content":[
    {
      "paragraph": "<strong>Go语言版:</strong>",
      "preCode":\"${code}\"
    }
  ]
}'

