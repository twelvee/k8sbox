# k8sbox
**Helpful tool to replicate your k8s environments**

### Run local build with docker
<code>
make docker-build<br>
docker run -v ${PWD}/examples/environments:/tmp/environments k8sbox:current run -f /tmp/environments/example_environment.toml
</code>