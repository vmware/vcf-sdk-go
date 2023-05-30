SWAGGER_VERSION=0.30.3

check-swagger:
	@swagger version | grep ${SWAGGER_VERSION} > /dev/null || { echo "Wrong version of swagger command. Install go-swagger ${SWAGGER_VERSION}"; exit 1; }

swagger: check-swagger
	rm -rf client models
	swagger generate client -r copyright.txt \
		--name=VcfClient \
		--additional-initialism=SDDC \
		--additional-initialism=NSXT \
		--additional-initialism=FIPS \
		--additional-initialism=CEIP \
		--additional-initialism=VSAN \
		--additional-initialism=VRA \
		--additional-initialism=VROPS \
		--additional-initialism=WSA \
		--additional-initialism=VCENTERS \
		--additional-initialism=SOS \
		--additional-initialism=SAN \
		--additional-initialism=VRSLCM \
		--additional-initialism=CSR \
		--additional-initialism=CSRS \
		--additional-initialism=SSO \
		--additional-initialism=ISO \
		--additional-initialism=SSL \
		--additional-initialism=VMWARE \
		-f ./swagger.json
	./scripts/fix_abbreviations.sh


modified:
	git ls-files --modified | xargs git add

build:
	go build