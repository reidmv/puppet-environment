class reidmv::packages {

  Package {
    ensure => installed,
  }

  package { 'git': }
  package { 'vim': }
  package { 'puppet':
    ensure => latest,
  }

}
