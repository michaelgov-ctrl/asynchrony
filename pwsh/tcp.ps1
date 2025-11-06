
class Listener {
    [System.Net.Sockets.TcpListener]$L
    [int]$Port

    Listener() {
        $this.Port = 4242
        $this.L = New-Object System.Net.Sockets.TcpListener(New-Object System.Net.IPEndPoint([System.Net.IPAddress]::Any, $this.Port))
        $this.L.Start()
    }

    [void]Accept() {
        $msg = [System.Text.Encoding]::ASCII.GetBytes("you connected to port {0}!" -f $this.Port)
        $end = (Get-Date).AddMinutes(1)

        while ((Get-Date) -lt $end) {
            if (-not $this.L.Pending()) {
                Start-Sleep -Milliseconds 200
                continue
            }

            $c = $this.L.AcceptTcpClient()
            $conn = $c.GetStream()
            $conn.Write($msg, 0, $msg.Length)
            $conn.Flush()
                    
            $conn.Dispose()
        }

        $this.Close()
    }

    [void]Close() {
        $this.L.Close()
    }
}

class Client {
    [System.Net.Sockets.TcpClient]$C
    [int]$Port

    Client() {
        $this.Port = 4242
        $this.C = New-Object System.Net.Sockets.TcpClient("127.0.0.1", $this.Port)
    }

    [void]Connect() {
        Start-Sleep -Milliseconds 500 # alittle cheating for the demo

        $conn = $this.C.GetStream()
        $buf = New-Object byte[] 1024
        $len = $conn.Read($buf, 0, $buf.Length)
        Write-Host ("connected with handshake: {0}" -f [System.Text.Encoding]::ASCII.GetString($buf, 0, $len))
        $conn.Dispose()
        $this.Close()
    }

    [void]Close() {
        $this.C.Close()
    }
}
