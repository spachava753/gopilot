output "private_key_openssh" {
  description = "Private key data in OpenSSH PEM (RFC 4716) format"
  value       = try(trimspace(module.key_pair.private_key_openssh), "")
  sensitive   = true
}

output "private_key_pem" {
  description = "Private key data in PEM (RFC 1421) format"
  value       = try(trimspace(module.key_pair.private_key_pem), "")
  sensitive   = true
}

output "public_key_openssh" {
  description = "The public key data in \"Authorized Keys\" format. This is populated only if the configured private key is supported: this includes all `RSA` and `ED25519` keys"
  value       = try(trimspace(module.key_pair.public_key_openssh), "")
}

output "public_key_pem" {
  description = "Public key data in PEM (RFC 1421) format"
  value       = try(trimspace(module.key_pair.public_key_pem), "")
}