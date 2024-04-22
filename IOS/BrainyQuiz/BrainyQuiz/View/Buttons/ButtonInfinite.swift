//
//  ButtonInfinite.swift
//  BrainyQuiz
//
//  Created by Horacio Cortez on 16/04/2024.
//

import Foundation
import SwiftUI

struct ButtonInfinite: View {
    var text: String
    var action: () -> Void
    
    
    var body: some View {
        Button(action: action){
            Text(text)
                .frame(maxWidth: .infinity).padding(.vertical, 10)
                .foregroundStyle(.white)
                .background(Color.clear)
                .clipShape(RoundedRectangle(cornerRadius: 10))
                .overlay(RoundedRectangle(cornerRadius: 10).stroke(Color.orange, lineWidth: 2))
        }
    }
    
}

