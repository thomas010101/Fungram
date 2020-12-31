# Fungram

Build library to run in the mobile


cd cmd/hd_wallet


#build ios library

gomobile bind -tags 'debug' -target=ios ./

#build android library

gomobile bind -tags 'debug' -target=android ./
