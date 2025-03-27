# grammars
Grammars is a Lexer and Parser that takes any kind of data as input, apply its rules and conditions, and generates an AST based on the request

Describe how the file is composed (header, blocks, constants, rules).


## Rule Definition Syntax
A **rule** defines a sequence of characters or bytes that can be used to compose a token.

---

### Rule Structure
Each rule consists of:

- A **rule name**
- A **colon (`:`)**
- The **characters or bytes it represents**
- A **semicolon (`;`)** to end the definition

---

### Rule Naming Convention
- Rule names **must be in uppercase**
- Only **letters and underscores (`_`)** are allowed

---

### Rule Formats
#### Character Rule (use double quotes)
Use when the rule represents **text characters**:

```text
HELLO: "hello";
```

#### Byte Rule (use square brackets)
Use when the rule represents **byte values**:

```text
BYE: [23, 45, 56];
```

## Constant Definition Syntax
### What is a Constant?
A **constant** is a reusable composition made from:
- One or more **rules**
- Or other **constants**
- Along with an optional **number of times each should repeat**

---

### Constant Naming
- A constant must begin with an **underscore (`_`)**
- Then it should follow **CamelCase** naming (e.g. `_myConstant`)

---

### How to Use Constants
#### Adding Rules or Constants
To include a **rule** or another **constant** in a constant definition:
- Prefix it with a **dot (`.`)**

#### Repetition (Occurrences)
To specify how many times a rule/constant should repeat:
- Add the number in **square brackets `[n]`**
- If no number is provided, it defaults to `1`

---

### Example
#### Final constant:
This constant matches the following character sequence twice:
```
--my-username@mydomain.com|my-username@mydomain.com|
```

#### Constant definition:
```text
myCharacters: .HYPHEN[2] ._myEmailWithPipe[2];
```

#### Referenced constant:
```text
_myEmailWithPipe: .USERNAME .COMMERCIAL_A .DOMAIN .DOT .COM .PIPE;
```

#### Rules used:
```text
PIPE: "|";
HYPHEN: "-";
USERNAME: "my-username";
COMMERCIAL_A: "@";
DOMAIN: "mydomain";
DOT: ".";
COM: "com";
```



## Token
### Reverse
### Escaping
### Reference
### Must be unique
### Must not be unique
### Selector
#### Is Not
### Cardinality



## Block
### Line
### Unit tests
#### Must match
#### Must not match

## Header
### Version
### Entry Point
### Omission