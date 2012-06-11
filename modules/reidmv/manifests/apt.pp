class reidmv::apt {

  class { '::apt':
    purge_sources_list   => true,
    purge_sources_list_d => true,
  }

  apt::source { "puppetlabs":
    include_src => true,
    key         => "4BD6EC30",
    key_server  => "pgp.mit.edu",
    location    => "http://apt.puppetlabs.com",
    release     => "${lsbdistcodename}",
    repos       => "main",
  }

  apt::source { 'mirrors.cat.pdx.edu':
    location    => "http://mirrors.cat.pdx.edu/ubuntu",
    repos       => "main restricted universe multiverse",
    key         => "437d05b5",
    key_server  => "pgp.mit.edu",
    include_src => true,
  }

  apt::source { 'mirrors.cat.pdx.edu updates':
    location    => "http://mirrors.cat.pdx.edu/ubuntu",
    repos       => "main restricted universe multiverse",
    key         => "437d05b5",
    key_server  => "pgp.mit.edu",
    include_src => true,
    release     => "${lsbdistcodename}-updates",
  }

  apt::source { 'mirrors.cat.pdx.edu security':
    location    => "http://mirrors.cat.pdx.edu/ubuntu",
    repos       => "main restricted universe multiverse",
    key         => "437d05b5",
    key_server  => "pgp.mit.edu",
    include_src => true,
    release     => "${lsbdistcodename}-security",
  }

}
