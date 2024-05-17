 # ListIME.ps1
   Get-WinUserLanguageList | Select-Object -ExpandProperty InputMethodTips | ForEach-Object { 
       $_.LayoutName 
   }