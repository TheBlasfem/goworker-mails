require 'resque'

class Mailer
  def self.get_args
    #complex operation
    ['sender@gmail.com', 'receiver@outlook.com']
  end
end

Resque.push('mail', 'class' => 'Mailer', 'args' => Mailer.get_args)