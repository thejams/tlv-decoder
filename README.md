# tlv-decoder

a simple golang tlv-decoder that receive a string with a TLV format and returns a map with the content of the TLV information

# compilation with docker

- make sure you have docker installed

- in the root of the proyect, build the image with the following command docker build -t tlv-decoder .

- then, run the container in a interactive way with the following command docker run --interactive tlv-decoder

- the console will ask you for a string to enter, this is where you sould enter the string with TLV format
