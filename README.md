# Tracking

This tool helps generating invoices from data stored in markdown files.

There are three different file for each kind of data; clients, hours and invoices.

Run `./tracking -h` to see the available options.

## Clients file

The template for the clients file is:

```markdown
# <CLIENT_KEY>
<BILLING_INFO>
```

For example:

```markdown
# foobar
Foo Bar Manager
Foo Bar Ltd (12345678)
foo@bar.io
12 Foo bar st, Foobarland
VAT: FB12345678
MOD: 98273311

# zap
ZAP Manager
ZAP Ltd (12345678)
zap@email.io
12 ZAP st, Foobarland
VAT: FB87654321
MOD: 8877665
```

## Hours file

The template for the hours file is:

```markdown
# <DAY>
<NUM> <CLIENT_KEY>
```

For example:

```markdown
# 2018-09-12
1 foobar
5 zap

# 2018-09-13
1 foobar
3 zap
```

## Invoices file

The template for the hours file is:
```markdown
<INVOICE_ID> <CLIENT_KEY> <HOURLY_RATE> <FROM_DATE> <TO_DATE>
```

For example:

```markdown
100001 foobar 12.5 2018-09-12 2018-09-13
```
