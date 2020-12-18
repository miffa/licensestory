

# private key
openssl genpkey -algorithm  RSA  -outform PEM -out private_key.pem

# publice key
openssl rsa -in private_key.pem -outform PEM -pubout -out public_key.pem


# encrypt
openssl rsautl -encrypt -inkey public_key.pem -pubin -in plain_text.log -out encrypted_text.log

#decrypt
openssl rsautl -decrypt -inkey private_key.pem -in encrypted_text.log -out decrypted_text.log




# encrypt uuid
openssl rsautl -encrypt -inkey  rsa_public.pem  -pubin -in uuid.log  -out uuid_encrypt.log && base64 uuid_encrypt.log > uuid_e_b.log



# decrypt uuid
base64 -d uuid_e_b.log > temp.log &&openssl rsautl -decrypt -inkey rsa_private.pem -in temp.log -out uuid_ok.log && rm -rf temp.log && cat uuid_ok.log


# sign
openssl dgst -sha256 -sign <private-key> -out /tmp/sign.sha256 <file>

# verify
openssl dgst -sha256 -verify <pub-key> -signature /tmp/sign.sha256 <file>


