//
//  WelcomeView.swift
//  BrainyQuiz
//
//  Created by Horacio Cortez on 16/04/2024.
//

import Foundation
import SwiftUI

struct WelcomeView: View {
    var body: some View {
        NavigationView{
            ZStack{
                Color.black.ignoresSafeArea(.all)
                
                VStack{
                    Spacer()
                    Image("BrainyQuiz").resizable().aspectRatio(contentMode: .fit).frame(width: 200, height: 200)
                    Text("Desafiate a ti mismo y a tus amigos!").font(.headline).foregroundStyle(.white)
                        .padding(.top, 10)
                    
                    Spacer()
                    VStack {
                     
                        NavigationLink(destination: RegisterView(), label: {
                            ButtonStartText(text: "Inicia Sesion")
                        })
                        .padding(.vertical, 10)
                        
                        NavigationLink(destination: RegisterView(), label: {
                            ButtonStartText(text: "Registrate")
                        })
                        DividerText(text: "o", color: Color.white).padding(.vertical)
                        Button(action: {
                            print("google")
                        }) {
                            HStack {
                                Image("Google").resizable().aspectRatio(contentMode: .fill).frame(width:20, height: 20)
                                    .font(.title)
                                Text("Inicia sesión con Google")
                                    .fontWeight(.bold).foregroundStyle(.white)
                            }
                            .frame(maxWidth: .infinity)
                            .padding() // Agrega espacio alrededor del contenido del botón
                            .background(Color.clear)
                            .clipShape(RoundedRectangle(cornerRadius: 10))
                            .overlay(RoundedRectangle(cornerRadius: 10).stroke(Color.orange, lineWidth: 2))
                          
                        }.frame(maxWidth: .infinity)
                    }.padding(.horizontal, 40).padding(.bottom, 40)
                    Spacer()
                    Text("@hackatondeldev 2024").font(.caption).foregroundStyle(.white)
                }
            }
        }
    }
};



                                       
                                       
#Preview {
    WelcomeView()
}


struct ButtonStartText: View {
    var text: String
    var body: some View {
        Text(text)
            .frame(maxWidth: .infinity)
            .foregroundStyle(.white)
            .padding()
            .clipShape(RoundedRectangle(cornerRadius: 10))
            .overlay(RoundedRectangle(cornerRadius: 10).stroke(Color.orange, lineWidth: 2))
    }
}

