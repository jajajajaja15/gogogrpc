GEN_PB_DIR := .

proto:
	protoc --go_out=${GEN_PB_DIR} \
	