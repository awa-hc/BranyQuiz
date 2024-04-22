//
//  RegisterView.swift
//  BrainyQuiz
//
//  Created by Horacio Cortez on 16/04/2024.
//

import Foundation
import SwiftUI


struct RegisterView: View {
    
    @State private var username = ""
    @State private var email = ""
    @State private var password = ""
    @State private var confirmPassword = ""
    
    @State private var showAlert = false
    @State private var errorMessage = ""
    
    var body: some View {
        ZStack{
            Color.black.ignoresSafeArea()
            
            VStack(alignment: .leading, spacing: 20) {
                Text("Registro")
                    .font(.title)
                    .foregroundStyle(.white)
                    .fontWeight(.bold)
                    .padding(.bottom, 20)
                FormField(text: $username, title: "Ingresa tu nombre de usuario", placeholder: "username")

                FormField(text: $email, title: "Ingresa tu email", placeholder: "email")
                
                FormField(text: $password, title: "Ingresa tu password", placeholder: "password")
                FormField(text: $confirmPassword, title: "Repite tu password", placeholder: "re -password")
                
                DividerText(text: "Enviar", color: Color.white)
                ButtonInfinite(text: "Registrarse"){
                    
                    if validateInput() {
                        print("registro exitoso")
                    } else {
                     showAlert = true
                    }
                }
            }
            .padding(20)
            .alert(isPresented: $showAlert){
                Alert(
                    title: Text("error"),
                    message: Text(errorMessage),
                    dismissButton: .default(Text("Ok")))
            }
            
        }
    }
    
    private func validateInput() -> Bool {
           let emailRegex = #"^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}$"#
           let usernameRegex = #"^[a-zA-Z0-9]{3,20}$"#
           let passwordRegex = #"^(?=.*[0-9])(?=.*[!@#$%^&*])(?=.*[a-z])(?=.*[A-Z]).{8,}$"#
           
           if !NSPredicate(format: "SELF MATCHES %@", emailRegex).evaluate(with: email) {
               errorMessage = "Email inválido"
               return false
           } else if !NSPredicate(format: "SELF MATCHES %@", usernameRegex).evaluate(with: username) {
               errorMessage = "Username debe tener entre 3 y 20 caracteres alfanuméricos"
               return false
           } else if !NSPredicate(format: "SELF MATCHES %@", passwordRegex).evaluate(with: password) {
               errorMessage = "Password debe contener al menos 1 número, 1 caracter especial, 1 mayúscula y tener mínimo 8 caracteres"
               return false
           } else if password != confirmPassword {
               errorMessage = "Las contraseñas no coinciden"
               return false
           }
           
           return true
       }
    
    
}


#Preview {
    RegisterView()
}

struct FormField: View {
    @Binding var text: String
    var title: String
    var placeholder: String
    var body: some View{
        Text(title).foregroundStyle(.white).padding(1)
        TextField(placeholder, text: $text)
            .padding(.vertical, 10)
            .padding(.horizontal, 15)
            .background(
                RoundedRectangle(cornerRadius: 10)
                    .stroke(Color.orange, lineWidth: 2)
                    .background(Color.white.opacity(0.2))
                    .overlay(
                        RoundedRectangle(cornerRadius: 10)
                            .stroke(Color.clear, lineWidth: 2)
                    )
            )
            .padding(1)
            .cornerRadius(10)
            .foregroundColor(.white)
            .font(.body)
    }
}
