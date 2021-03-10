# gene-beauty-ether-addr

[![view examples](https://img.shields.io/badge/learn%20by-examples-0C8EC5.svg?style=for-the-badge&logo=go)](https://github.com/pefish/gene-beauty-ether-addr)

Read this in other languages: [English](README.md), [简体中文](README_zh-cn.md)

gene-beauty-ether-addr

## Install

```
go get github.com/pefish/gene-beauty-ether-addr/cmd/...
```

## Quick start

```shell script
gene-beauty-ether-addr --config=/path/to/config
```

or

```shell script
GO_CONFIG=/path/to/config gene-beauty-ether-addr
```

## Db

```sql
CREATE TABLE `address` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `address` varchar(255) NOT NULL,
  `path` varchar(255) NOT NULL COMMENT 'path',
  `mnemonic` varchar(255) NOT NULL COMMENT 'mnemonic',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `address` (`address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

```

## Document

[doc](https://godoc.org/github.com/pefish/gene-beauty-ether-addr)

## Contributing

1. Fork it
2. Download your fork to your PC
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Make changes and add them (`git add .`)
5. Commit your changes (`git commit -m 'Add some feature'`)
6. Push to the branch (`git push origin my-new-feature`)
7. Create new pull request

## Security Vulnerabilities

If you discover a security vulnerability, please send an e-mail to [pefish@qq.com](mailto:pefish@qq.com). All security vulnerabilities will be promptly addressed.

## License

This project is licensed under the [Apache License](LICENSE).
