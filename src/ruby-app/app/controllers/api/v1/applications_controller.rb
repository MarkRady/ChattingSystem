class Api::V1::ApplicationsController < ApplicationController

    def index
        applications = Application.select(:Token, :Name).order("Id Desc")
        render json: {data: applications, chatsInCach: chatsInCach}, status: :ok
    end

    def create
        application = Application.new(application_params)
        if application.save 
            render json: {data: {
                Name: application.Name,
                Token: application.Token
            }}, status: :ok
        else
            render json: {data: application.errors}, status: :unprocessable_entity
        end
    end

    def update
        application = Application.find_by(Token: params[:id])
        if application.update_attributes(application_params)
            render json: {data: {
                Name: application.Name,
                Token: application.Token
            }}, status: :ok
        else
            render json: {data: application.errors}, status: :unprocessable_entity
        end
    end

    
    def show
        application = Application.select(:Token, :Name).find_by(Token: params[:id])
        render json: {data: application}, status: :ok
    end

    

    private
    
    def application_params
        params.permit(:Name)
    end

end