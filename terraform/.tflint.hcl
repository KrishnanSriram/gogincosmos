# Require that all providers have version constraints through required_providers.
rule "terraform_required_providers"{
    enabled = true
}

# Disallow // comments in favor of #
rule "terraform_comment_syntax" {
    enabled = false
}

# terraform_documented_outputs
rule "terraform_documented_outputs" {
    enabled = true
}

# terraform_documented_variables
rule "terraform_documented_variables" {
    enabled = true
}