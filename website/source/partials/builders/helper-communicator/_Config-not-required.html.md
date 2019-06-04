<!-- Code generated from the comments of the Config struct in helper/communicator/config.go; DO NOT EDIT MANUALLY -->

-   `communicator` (string) - Type
-   `ssh_host` (string) - SSH
    
-   `ssh_port` (int) - SSH Port
-   `ssh_username` (string) - SSH Username
-   `ssh_password` (string) - SSH Password
-   `ssh_keypair_name` (string) - If specified, this is the key that will be used for SSH with the
    machine. The key must match a key pair name loaded up into Amazon EC2.
    By default, this is blank, and Packer will generate a temporary keypair
    unless [`ssh_password`](../templates/communicator.html#ssh_password) is
    used.
    [`ssh_private_key_file`](../templates/communicator.html#ssh_private_key_file)
    or `ssh_agent_auth` must be specified when `ssh_keypair_name` is
    utilized.
    
-   `temporary_key_pair_name` (string) - SSH Temporary Key Pair Name
-   `ssh_clear_authorized_keys` (bool) - SSH Clear Authorized Keys
-   `ssh_private_key_file` (string) - SSH Private Key File
-   `ssh_interface` (string) - SSH Interface
-   `ssh_ip_version` (string) - SSHIP Version
-   `ssh_pty` (bool) - SSH Pty
-   `ssh_timeout` (time.Duration) - SSH Timeout
-   `ssh_agent_auth` (bool) - If true, the local SSH agent will be used to authenticate connections to
    the source instance. No temporary keypair will be created, and the
    values of `ssh_password` and `ssh_private_key_file` will be ignored. To
    use this option with a key pair already configured in the source AMI,
    leave the `ssh_keypair_name` blank. To associate an existing key pair in
    AWS with the source instance, set the `ssh_keypair_name` field to the
    name of the key pair.
    
-   `ssh_disable_agent_forwarding` (bool) - SSH Disable Agent Forwarding
-   `ssh_handshake_attempts` (int) - SSH Handshake Attempts
-   `ssh_bastion_host` (string) - SSH Bastion Host
-   `ssh_bastion_port` (int) - SSH Bastion Port
-   `ssh_bastion_agent_auth` (bool) - SSH Bastion Agent Auth
-   `ssh_bastion_username` (string) - SSH Bastion Username
-   `ssh_bastion_password` (string) - SSH Bastion Password
-   `ssh_bastion_private_key_file` (string) - SSH Bastion Private Key File
-   `ssh_file_transfer_method` (string) - SSH File Transfer Method
-   `ssh_proxy_host` (string) - SSH Proxy Host
-   `ssh_proxy_port` (int) - SSH Proxy Port
-   `ssh_proxy_username` (string) - SSH Proxy Username
-   `ssh_proxy_password` (string) - SSH Proxy Password
-   `ssh_keep_alive_interval` (time.Duration) - SSH Keep Alive Interval
-   `ssh_read_write_timeout` (time.Duration) - SSH Read Write Timeout
-   `winrm_username` (string) - WinRM
    
-   `winrm_password` (string) - Win RM Password
-   `winrm_host` (string) - Win RM Host
-   `winrm_port` (int) - Win RM Port
-   `winrm_timeout` (time.Duration) - Win RM Timeout
-   `winrm_use_ssl` (bool) - Win RM Use SSL
-   `winrm_insecure` (bool) - Win RM Insecure
-   `winrm_use_ntlm` (bool) - Win RM Use NTLM
-   `pause_before_connecting` (time.Duration) - Delay
    