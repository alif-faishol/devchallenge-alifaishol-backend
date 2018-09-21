# Back-End Aplikasi Telkom Codex

- **Nama tim**: alifaishol
- **Anggota tim**: Muhammad Alif Faishol

## Cara Menjalankan Aplikasi

1. Perlu `go`.

2. Masuk ke $GOPATH, jika belum:
```
$ export GOPATH=$(go env GOPATH)
$ cd $GOPATH
```

2. Download repo dan *dependency*:

```
$ go get https://github.com/alif-faishol/devchallenge-alifaishol-backend
```

3. Masuk ke folder repo:
```
$ cd src/github.com/alif-faishol/devchallenge-alifaishol-backend
```

4. Copy `.env.dist` ke `.env`.
```
$ cp .env.dist .env
```

5. Pastikan alamat aplikasi front-end di `.env` sudah sesuai.

6. Jalankan development server:
```
go run *.go
```
