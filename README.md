<div align="center">
    <h1>passctl | server</h1>
    <h3>Passctl web-server written in go</h3>
</div>

---

### Installation (Docker)
1. Clone this repository:
```bash
git clone https://github.com/passctl/server.git && cd server 
```
2. Copy over the default configuration
```
cp config.json.example config.json
```
3. Build and start the docker 
```
docker build -t passctl .
docker run -d -v $PWD/config.json:/app/config.json \
              -v $PWD/db:/app/db \
              -p [local port]:8080 \
              --restart unless-stopped \
              passctl 
```

### Configuration 
Server configuration can be edited using the `config.json` file. Note that you will 
need to restart the server after configuration changes.

#### `port`
Interface and port that the server will listen on. 

#### `password`
Password that will be asked for vault generation. Leave empty if you want to 
disable it.

> ⚠️ Please do not disable key generation password on public servers as it 
would allow mass key generation which may result in a DoS attack.

#### `max_vault`
Maximum amount of vaults allowed for generation. Set to `0` if you want no 
limitation.

#### `max_vault_size`
Maximum vault size (in megabytes). Set to `0` if you want no limitation. 
