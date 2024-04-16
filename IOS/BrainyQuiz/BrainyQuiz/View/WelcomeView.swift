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
        VStack {
            Text("!Bienvenido a tu juego de trivia!").font(.title).padding()
            
            Button("Iniciar Sesion") {
                
            }
            .padding().foregroundColor(.white)
            .background(Color.blue).clipShape(RoundedRectangle(cornerRadius: 10))
        }.padding()
    }
}


#Preview {
    WelcomeView()
}
