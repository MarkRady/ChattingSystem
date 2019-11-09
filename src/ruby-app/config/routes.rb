require 'sidekiq/web'
require 'sidekiq-scheduler/web'

Rails.application.routes.draw do
  mount Sidekiq::Web => '/sidekiq'

  namespace 'api' do 
    namespace 'v1' do 
      resources :applications
    end
  end
end
