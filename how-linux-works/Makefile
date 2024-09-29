.PHONY: create-image
create-image:
	docker image build -t linux-in-practice-2nd-docker .

.PHONEY: into-container
into-container:
	docker run --rm -it linux-in-practice-2nd-docker /bin/bash