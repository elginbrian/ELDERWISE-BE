$projectRoot = "d:\ELDERWISE-BE"
$goFiles = Get-ChildItem -Path $projectRoot -Filter "*.go" -Recurse

# Function to check if a comment is important and should be preserved
function Is-ImportantComment {
    param (
        [string]$commentText
    )
    
    # List of patterns for important comments
    $patterns = @(
        'godoc',
        'swagger:',
        '@\w+',
        '\+\w+',
        'TODO:',
        'FIXME:',
        'NOTE:'
    )
    
    # Check if comment matches any important pattern
    foreach ($pattern in $patterns) {
        if ($commentText -match $pattern) {
            return $true
        }
    }
    
    return $false
}

# Process each Go file
foreach ($file in $goFiles) {
    Write-Host "Processing: $($file.FullName)"
    
    # Read file content
    $content = Get-Content -Path $file.FullName -Raw
    $lines = $content -split "`r`n|\r|\n"
    $newLines = @()
    
    # Process each line
    foreach ($line in $lines) {
        # Skip empty lines or add them as is
        if ([string]::IsNullOrWhiteSpace($line)) {
            $newLines += $line
            continue
        }
        
        # Handle full-line comments first
        if ($line -match '^\s*//') {
            # Check if it's an important comment
            if (Is-ImportantComment -commentText $line) {
                $newLines += $line
            }
            continue
        }
        
        # Find // comment position, considering strings
        $commentPos = -1
        $inString = $false
        $stringChar = ''
        $escape = $false
        
        for ($i = 0; $i -lt $line.Length - 1; $i++) {
            $char = $line[$i]
            $nextChar = $line[$i + 1]
            
            # Handle string detection
            if (($char -eq '"' -or $char -eq "'") -and -not $escape) {
                if (-not $inString) {
                    $inString = $true
                    $stringChar = $char
                } 
                elseif ($char -eq $stringChar) {
                    $inString = $false
                }
            }
            
            # Handle escaped characters
            $escape = ($char -eq '\' -and -not $escape)
            
            # Detect comment start outside of strings
            if (-not $inString -and $char -eq '/' -and $nextChar -eq '/') {
                $commentPos = $i
                break
            }
        }
        
        # No comment found, add the whole line
        if ($commentPos -eq -1) {
            $newLines += $line
            continue
        }
        
        # Get the comment part
        $commentPart = $line.Substring($commentPos)
        $codePart = $line.Substring(0, $commentPos).TrimEnd()
        
        # Check for URLs in comment (to avoid treating http:// as a comment)
        $hasUrl = $commentPart -match 'https?://'
        
        # Check if comment is important
        if ($hasUrl -or (Is-ImportantComment -commentText $commentPart)) {
            $newLines += $line
        } 
        else {
            # Add only the code part without the comment
            $newLines += $codePart
        }
    }
    Set-Content -Path $file.FullName -Value ($newLines -join [Environment]::NewLine)
} # End of foreach ($file in $goFiles)

Write-Host "Comment removal completed! Important comments were preserved."
