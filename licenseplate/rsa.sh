
# private key
openssl genpkey -algorithm  RSA  -outform PEM -out private_key.pem

# publice key
openssl rsa -in private_key.pem -outform PEM -pubout -out public_key.pem


# encrypt
openssl rsautl -encrypt -inkey public_key.pem -pubin -in plain_text.log -out encrypted_text.log

#decrypt
openssl rsautl -decrypt -inkey private_key.pem -in encrypted_text.log -out decrypted_text.log
