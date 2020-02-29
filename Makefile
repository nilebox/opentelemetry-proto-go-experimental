GENDIR := gen

# Find all .proto files.
PROTO_FILES := $(wildcard proto/opentelemetry/proto/*/v1/*.proto proto/opentelemetry/proto/collector/*/v1/*.proto)

# Function to execute a command. Note the empty line before endef to make sure each command
# gets executed separately instead of concatenated with previous one.
# Accepts command to execute as first parameter.
define exec-command
$(1)

endef

# Generate ProtoBuf implementation for Go.
.PHONY: gen-go
gen-go:
	$(foreach file,$(PROTO_FILES),$(call exec-command,protoc --go_out=plugins=grpc:$(GOPATH)/src $(file)))
	rm -rf ./$(GENDIR)/go
	cp -R $(GOPATH)/src/github.com/open-telemetry/opentelemetry-proto/$(GENDIR)/go ./gen/