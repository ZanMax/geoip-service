# GeoIP Service

GeoIP Service to detect the country and city of the client

### Install

```bash
mkdir geoip-data
```
Copy the GeoLite2-City.mmdb and GeoLite2-Country.mmdb file into the geoip-data folder

### Usage

http://localhost:8000/geoip/city/<IP>
http://localhost:8000/geoip/country/<IP>

### GeoIP Update

```bash
sudo add-apt-repository ppa:maxmind/ppa
sudo apt-get update
sudo apt-get install geoipupdate
```

### Provider

- MaxMind GeoIP

