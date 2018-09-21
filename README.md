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

2. Download repo:

```
$ go get github.com/alif-faishol/devchallenge-alifaishol-backend
```

3. Masuk ke folder repo:
```
$ cd src/github.com/alif-faishol/devchallenge-alifaishol-backend
```

4. Download *dependency*:
```
$ go get
```

5. Copy `.env.dist` ke `.env`.
```
$ cp .env.dist .env
```

6. Pastikan alamat aplikasi front-end di `.env` sudah sesuai.

7. Jalankan development server:
```
go run *.go
```
