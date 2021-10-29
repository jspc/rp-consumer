TYPES_DIR := types
GENERATED := $(TYPES_DIR)/location.go \
	     $(TYPES_DIR)/location.custom.go \
	     $(TYPES_DIR)/precipitation.go \
	     $(TYPES_DIR)/precipitation.custom.go \
	     $(TYPES_DIR)/sensorreading.go \
	     $(TYPES_DIR)/sensorreading.custom.go \
	     $(TYPES_DIR)/weatherforecast.go \
	     $(TYPES_DIR)/weatherforecast.custom.go

default: $(GENERATED)

$(GENERATED): mint/*.mint
	mint generate mint/ -f
