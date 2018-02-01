## Simple file reveiver script

### Server side setup

1. down the complied binary [filereceiver](https://github.com/RHVH-QE/file_receiver/releases/download/v0.1.0/filereceiver), and `chmod +x`

2. edit the `.env` file, and keep this file at same level with the script

* `ROOT_PATH` is where should the file uploaded file to save, must be absolut path
* `PORT` is the listen port, this will be effected the `curl` command, make make sure the `port` not blocked by the firewall

### Client side usage

1. upload single file to the `ROOT_PATH`

```
curl -X POST http://{server_ip}:{port}/upload \
  -F "upload[]=@/Users/appleboy/test1.zip" \
  -H "Content-Type: multipart/form-data"
```

2. upload multiple file to the sub directory under `ROOT_PATH`

```
curl -X POST http://{server_ip}:{port}/upload/{arbitary_dir} \
  -F "upload[]=@/Users/appleboy/test1.zip" \
  -F "upload[]=@/Users/appleboy/test2.zip" \
  -H "Content-Type: multipart/form-data"
```

two files will upload to `/{ROOT_PATH}/{arbitary_dir}`, if the `arbitary_dir` not exists, then will be created for u
