class reidmv::packages {
  require reidmv::apt

  Package {
    ensure => installed,
  }

  package { 'git': }
  package { 'vim': }
  package { 'puppet':
    ensure => latest,
  }

}
