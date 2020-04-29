# SQL Server CLI utility

This CLI provides some quick shortcuts for frequent workflows by wrapping Microsoft's sqlcmd.

## Commands 

### Query

Run arbitrary sql statements.

	sqlcli query "SET NOCOUNT ON; select top 2 IdCuenta, Carpeta from accounts.cuenta"

```
IdCuenta                             Carpeta                                                                         
------------------------------------ --------------------------------------------------------------------------------
52893024-A0F5-4F30-A79E-00006BD75EC9 NULL                                                                            
9A76B2C2-84A1-4D87-9511-00008BE34381 NULL                                                                            
```

### Script

Run statements from a .sql (batch) file.

	sqlcli script /tmp/insert-data.sql 

```
(89 rows affected)
```

### Find

Search objects (tables, views, stored procs or functions) containing text.

	sqlcli find "branding"

```
Object Name                                             Object Type          TEXT Location     
------------------------------------------------------- -------------------- ------------------
accounts.CuentaBranding                                 BASE TABLE           Table Name        
accounts.CuentaBranding.IdCuentaBranding                COLUMN               Column Name       
accounts.CuentaBranding_TMP                             BASE TABLE           Table Name        
accounts.CuentaBranding_TMP.IdCuentaBranding            COLUMN               Column Name       
accounts.DeleteCuentaBranding                           PROCEDURE            ROUTINE_DEFINITION
accounts.GetAllCuentaBranding                           PROCEDURE            ROUTINE_DEFINITION
accounts.GetCuentaBrandingByIdCuenta                    PROCEDURE            ROUTINE_DEFINITION
accounts.GetCuentaBrandingBySubDomain                   PROCEDURE            ROUTINE_DEFINITION
accounts.InsertCuentaBranding                           PROCEDURE            ROUTINE_DEFINITION
accounts.UpdateCuentaBranding                           PROCEDURE            ROUTINE_DEFINITION
reception.getSuscripcionesByIdCuenta                    PROCEDURE            ROUTINE_DEFINITION
reception.GetVentanillasNoSuscritas                     PROCEDURE            ROUTINE_DEFINITION
selfBilling.Branding                                    BASE TABLE           Table Name        
selfBilling.deleteBrandingById                          PROCEDURE            ROUTINE_DEFINITION
selfBilling.getBrandingById                             PROCEDURE            ROUTINE_DEFINITION
selfBilling.upsertBrandingById                          PROCEDURE            ROUTINE_DEFINITION

(16 rows affected)
```
### Get

Outputs the definition of a table, view, procedure or function

	sqlcli get selfBilling.getBrandingById

```
CREATE PROCEDURE selfBilling.getBrandingById
(@idCuenta uniqueidentifier)
AS

SELECT 
	IdPrimaryImage   ,
	IdSecondaryImage ,
	IdBackgroundImage,
	PrimaryColor     ,
	SecondaryColor   ,
	BackgroundColor  ,
	FontColor
FROM 
	selfBilling.Branding
WHERE
	IdCuenta = @idCuenta
```

## Installation

Download and install [SqlCmd](https://docs.microsoft.com/en-us/sql/tools/sqlcmd-utility?view=sql-server-ver15).

Compile (Go 1.14) and copy to `/usr/local/bin`:

	make install

## Configuration

Username and password authentication supported. (Integrated security support can be easily added)

	cat ~/.sqlcli.yaml 

```
SqlCmdBinary: /usr/local/bin/sqlcmd
Server: <ip address>
Database: <default database>
Username: <myusername>
Password: <mypassword>
```

