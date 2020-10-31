package awsdata

import "github.com/manywho/awsinventory/internal/inventory"

// ProcessRow takes an inventory row, performs some action, and returns an error
type ProcessRow func(inventory.Row) error
