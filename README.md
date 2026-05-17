# chemcli

CLI tool for managing pharmacy/medicine stock.

## Commands

### Add Medicine
Adds new medicine to the stock or updates existing stock quantity.
```bash
chemcli add -n <name> -q <quantity> -e <expiry_date>
```
* `-n, --name`: Medicine name (required)
* `-q, --qty`: Medicine quantity (required)
* `-e, --expiry`: Medicine expiry date in YYYY-MM-DD format (optional)

### List Stock
Lists all available medicines in the database.
```bash
chemcli list
```

### Search Medicine
Searches for a medicine by name.
```bash
chemcli search <medicine_name>
```

### Sell Medicine
Removes the specified quantity of a medicine from the stock.
```bash
chemcli sell -n <name> -q <quantity>
```
* `-n, --name`: Medicine name
* `-q, --qty`: Quantity to sell

### Low Stock Alert
Lists medicines that have a quantity at or below a specified threshold.
```bash
chemcli lowStock -t <threshold>
```
* `-t, --threshold`: Quantity threshold to check against
