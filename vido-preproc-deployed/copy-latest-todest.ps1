$destDirZip = "D:\Projects\GItHub\InvidositeHtmlgit\zips"
$dirZip = ".\VidoPreProc*.zip"

$latestZipPkg = Get-ChildItem -Path $dirZip | Sort-Object Name -Descending | Select-Object -First 1

#Write-Debug($latestZipPkg)
if ([string]::IsNullOrEmpty($latestZipPkg)) {
    Write-Host("No package found in directory $dirZip `nSorry I can't continue.")
    return
}

if ([string]::IsNullOrEmpty($destDirZip)) {
    Write-Host("Destination dir is empty.`nSorry I can't continue.")
    return
}

$confirmation = Read-Host "Copy $latestZipPkg to $destDirZip ?`nContinue (y/n)?"
if ($confirmation -ne 'y') {
  Write-Host "As you like, nothing to do here. Bye."
  return
}

Write-Host "copy the $latestZipPkg to $destDirZip" 
Copy-Item $latestZipPkg -Destination $destDirZip
Write-Host "Zip copied OK"
