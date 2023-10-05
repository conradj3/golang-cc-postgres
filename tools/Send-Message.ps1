function GetRandomString {
    $length = 8
    $characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    
    $random = New-Object System.Random
    $randomString = -join (0..($length - 1) | ForEach-Object { $characters[$random.Next(0, $characters.Length)] })
    
    return $randomString
}

function EnqueueRandomMessages($url, $count) {
    for ($i = 0; $i -lt $count; $i++) {
        $message = GetRandomString
        $body = @{ message = $message } | ConvertTo-Json
        Invoke-RestMethod -Uri $url -Method POST -Body $body -ContentType "application/json"
        Write-Host "Enqueued message: $message"
    }
}

$Url = "http://localhost:8080/enqueue"
$Count = 10

EnqueueRandomMessages $Url $Count
