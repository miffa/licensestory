#pFPr2A3stgRSH0Lqf8L1QeMQiy0XsJIVOVLhP8PL65Q=
tpaaslicense sign -f troila_license_test_yesexpire_80.pem -u pFPr2A3stgRSH0Lqf8L1QeMQiy0XsJIVOVLhP8PL65Q= -c '''{"corporation":"我们限额测试环境","quota":80, "expired_time":"2020-09-03 23:59:59","extension":"this is a test","version":"v2.2", "service_period":"一年", "home_license":"正式版"}'''  -p rsa

tpaaslicense sign -f troila_license_test_noexpire_80.pem -u pFPr2A3stgRSH0Lqf8L1QeMQiy0XsJIVOVLhP8PL65Q= -c '''{"corporation":"我们限额测试环境","quota":80, "expired_time":"无","extension":"this is a test","version":"v2.2", "service_period":"永久", "home_license":"正式版"}'''  -p rsa

tpaaslicense sign -f troila_license_test_noexpire_0.pem -u pFPr2A3stgRSH0Lqf8L1QeMQiy0XsJIVOVLhP8PL65Q= -c '''{"corporation":"我们限额测试环境","quota":0, "expired_time":"无","extension":"this is a test","version":"v2.2", "service_period":"永久", "home_license":"正式版"}'''  -p rsa
