### Crypto 加密库

#### usage
```
加解密算法在golang源码的 crypto/crypto.go
```

#### 常用的加解密：
- MD5
- AES对称加密
    - 加密模式
        - CBC (Cipher Block Chaining 密码分组链接模式)
        - ECB (Electronic Codebook 电码本模式)
        - CFB (Cipher Feedback 密码反馈模式)    
        - CTR (Counter 计算器模式)    
        - OFB (Output Feedback输出反馈模式)
    - 填充方式
        - PKCS5
        - PKCS7
        - ZERO
    - 输出格式
        - base64
        - hex