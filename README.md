# CLI Inventory Management System

## Overview

A simple command-line tool for managing product inventory. Users can add, view, update, and delete products, as well as generate CSV reports and import data from CSV files. The inventory state is saved to a JSON file for persistence.

## Preview

![Go CLI Inventory](https://dl.dropboxusercontent.com/scl/fi/hxumrpoyhhx303mri0q1o/go-cli-inventory.gif?rlkey=zzjotiuh33vt12a84ww3t934g&st=ybdw88wl&dl=1)

## Features

* **CRUD Operations:** Add, view, update, delete products.
* **CSV Import/Export:** Import data from and export reports to CSV.
* **Persistent Storage:** Saves/loads data using JSON.
* **Data Validation:** Checks input validity.
* **Error Handling:** Robust error management.
* **Duplicate Handling:** Option to keep or update existing entries on CSV import.
* **CLI Interface:** User-friendly command-line interaction.

## Installation

1.  **Build the application:**

    ```bash
    go build inventory.go operations.go data.go report.go
    ```

2.  **Run the executable:**

    ```bash
    ./inventory  # For macOS/Linux
    inventory.exe # For Windows
    ```

## CSV Import Format

```csv
Name,Category,Quantity,Price
Laptop,Electronics,10,1200.50
Mouse,Electronics,50,15.99
...