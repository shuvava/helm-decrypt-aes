# Helm decrypt AES data
decrypt logic for HELM encryptAES

* [encryptAES](https://helm.sh/docs/chart_template_guide/function_list/#encryptaes)

    ```kubernetes helm
    encryptAES "secretkey" "plaintext"
    # result below
    ```

* [decryptAES](https://helm.sh/docs/chart_template_guide/function_list/#decryptaes)

    ```kubernetes helm
    "30tEfhuJSVRhpG97XCuWgz2okj7L8vQ1s6V9zVUPeDQ=" | decryptAES "secretkey"
    # result 'plaintext'
    ```

This project has `DecryptAES` func allows to decrypt data encrypted by Helm