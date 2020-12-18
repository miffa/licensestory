#    生成rsa密钥对
##    默认名称   rsa_private.pem rsa_public.pem
##    [-p 密钥对前缀, 默认rsas, 可以包含地址]
./tpaaslicense genkey [-p xxxx]

#    使用以上生成的key生成license
##    -f license文件名
##    -u  企业ID（tpaas生成）
##    -c  json格式化的企业license信息
##    [-p 密钥对前缀, 可以包含地址]
./tpaaslicense sign -f testlicense.txt -u troila.com -c '''{"corporation":"troila.com","quota":160,"expired_time":"2022-02-03 23:59:59","extension":"this is a test","version":"v2.34", "service_period":"5 year"}''' [ -p xxxx ]

#    认证license文件信息
##    -f license文件
##    -u 企业ID
##    [-p 密钥对前缀, 可以包含地址]
./tpaaslicense verify -f testlicense.txt -u troila.com [ -p xxxx ]

# prefix 默认是rsa, 可以指定为字符串 或者 带路径的字符串
##    如果使用指定prefix -p aa
##        生成的名字是  aa_private.pem aa_public.pem
##  如果使用指定prefix -p  /data/license/me
##        生成的名字是  /data/license/me_private.pem /data/license/me_public.pem

#    解密uuid
## -u 加密之后的uuid
## -p 密钥前缀
tpaaslicense getuuid -p rsa -u 'VoMXf3L3b06OXzfPWsdVBkdNxm5g71oXjbri4OZKxQ79WPRpy07Uviye2+f5yu7tcMUPA+F8c/mQ
hdzipYU9pnnWV6PJWinE+Z/Ebe47L/d6D1xuE0tZ4SHRGCyp8Bkb6pCzcfniSBwzAtF8QwAz74KG
vQhhT5Rt43z6uBdE9U4='
