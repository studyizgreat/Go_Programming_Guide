
| Data Type       | Format Specifier | Description                                       |
|-----------------|------------------|---------------------------------------------------|
| **General**     | `%v`             | Default format, suitable for most types           |
|                 | `%+v`            | Adds field names for structs                      |
|                 | `%#v`            | Go-syntax representation of the value             |
|                 | `%T`             | Type of the value                                 |
|                 | `%%`             | Literal percent sign                              |
| **Boolean**     | `%t`             | Boolean: true or false                            |
| **Integer**     | `%d`             | Base 10                                           |
|                 | `%b`             | Base 2 (binary)                                   |
|                 | `%o`             | Base 8 (octal)                                    |
|                 | `%x`             | Base 16 (lowercase hexadecimal)                   |
|                 | `%X`             | Base 16 (uppercase hexadecimal)                   |
|                 | `%c`             | Character represented by the corresponding Unicode code point |
| **Floating Point and Complex** | `%f`   | Decimal point but no exponent, e.g., 123.456     |
|                 | `%e`             | Scientific notation (e.g., -1234.456e+78)         |
|                 | `%E`             | Scientific notation (e.g., -1234.456E+78)         |
|                 | `%g`             | Exponent as needed, only necessary digits         |
|                 | `%G`             | Exponent as needed, only necessary digits (uppercase) |
| **String and Byte Slice** | `%s`   | Plain string or slice of bytes                    |
|                 | `%q`             | Double-quoted string with Go syntax               |
|                 | `%x`             | Base 16, with two characters per byte             |
|                 | `%X`             | Base 16, with two characters per byte (uppercase) |
| **Pointer**     | `%p`             | Base 16 notation, with leading 0x                 |

