class reidmv::packages {
  require reidmv::apt

  Package {
    ensure => installed,
  }

  package { 'chromium-browser': ensure => latest; }
  package { 'facter': ensure => latest; }
  package { 'git': }
  package { 'puppet': ensure => latest; }
  package { 'ubuntu-restricted-extras': }
  package { 'vim': }

}
