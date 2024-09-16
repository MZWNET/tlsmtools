# tlsmtools - TL文件分割与合并工具

[English](README.md) | 简体中文

`tlsmtools` 是一个简单的命令行工具，用于分割和合并TL文件。

## 功能

- **分割**：将一个`tl`文件分割为`orig`和`alt`部分。
- **合并**：将`orig`、`alt`和`text`文件合并为一个`tl`文件。

## 安装

确保系统中已安装Go。目前仅在Go版本`1.23.1`上测试。

```bash
git clone https://github.com/MZWNET/tlsmtools.git
cd tlsmtools
go build
```

## 使用方法

### 分割文件

```bash
tlsmtools split -f [tl文件] [-d [输出目录]]
```

- `-f` 或 `--file`：需要分割的`tl`文件（必填）。
- `-d` 或 `--outputDir`：保存分割文件的目录（默认为当前目录）。

### 合并文件

```bash
tlsmtools merge -orig [orig文件] -alt [alt文件] -text [text文件] [-o [输出文件]]
```

- `-orig`：`orig`文件路径（必填）。
- `-alt`：`alt`文件路径（必填）。
- `-text`：`text`文件路径（必填）。
- `-o` 或 `--output`：输出路径及文件名（默认为 `output.tl`）。

## 许可证

本项目使用 [The Unlicense](http://unlicense.org/) 许可证。
