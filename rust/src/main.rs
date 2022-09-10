use wmi::*;
use std::collections::HashMap;
use wmi::Variant;

fn main() {
    let wmi_con = WMIConnection::new(COMLibrary::new()?.into())?;
    let results: Vec<HashMap<String, Variant>> = wmi_con.raw_query("SELECT * FROM Win32_OperatingSystem").unwrap();
    for os in results {
        println!("{:#?}", os);
    }
}
