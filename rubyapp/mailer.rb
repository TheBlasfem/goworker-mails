require 'resque'

class Mailer
  def self.get_args
    #complex operation
    ['receiver@outlook.com', 'Julio']
  end
end

Resque.push('mail', 'class' => 'Mailer', 'args' => Mailer.get_args)